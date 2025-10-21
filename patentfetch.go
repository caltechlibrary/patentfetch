package patentfetch

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

// PatentRecord is a simple struct to hold the information from each row after the header.
type PatentRecord struct {
	ID            string `csv:"id" json:"id,omitempty" yaml:"id,omitempty"`
	Title         string `csv:"title" json:"title,omitempty" yaml:"title,omitempty"`
	Assignee      string `csv:"assignee" json:"assignee,omitempty" yaml:"assignee,omitempty"`
	Inventor      string `csv:"inventor/author" json:"inventor_author,omitempty" yaml:"inventor_author,omitempty"`
	Priority      string `csv:"priority date" json:"priority_date,omitempty" yaml:"priority_date,omitempty"`            // Assuming this is a string representation of the date
	Filing        string `csv:"filing/creation date" json:"filing_creation,omitempty" yaml:"filing_creation,omitempty"` // Also a string representation
	Publication   string `csv:"publication date" json:"publication_date,omitempty" yaml:"publication_date,omitempty"`   // Also a string representation
	Grant         string `csv:"grant date" json:"grant_date,omitempty" yaml:"grant_date,omitempty"`                     // Also a string representation
	ResultLink    string `csv:"result link" json:"result_link,omitempty" yaml:"result_link,omitempty"`
	FigureLink    string `csv:"representative figure link" json:"representative_figure_link,omitempty" yaml:"representative_figure_link,omitempty"`
	DownloadedPDF string `csv:"downloaded pdf" json:"downloaded_pdf,omitempty" yaml:"downlaoded_pdf,omitempty"`
}

// Parse takes a CSV source text and parses the CSV file to return a list of Patent records from the CSV data
func Parse(src []byte) ([]*PatentRecord, error) {
	// NOTE: Trim first line, it's not a header, then we can parse it using the csv package
	lines := strings.Split(fmt.Sprintf("%s", src), "\n")
	txt := strings.Join(lines[1:], "\n")

	r := csv.NewReader(strings.NewReader(txt))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	data := []*PatentRecord{}
	for i, row := range records {
		fmt.Printf("DEBUG row %d -> %+v\n", i, row)
		/*
		           // Parse the single field into a PatentRecord
		           p := &PatentRecord{}

		           // If you want to parse the dates as time.Time objects, uncomment and adjust the following lines
		           // p.ParseDates()

		           // Print the parsed data (or process it further)
		           fmt.Printf("%+v\n", p)
		   		data = append(data, p)
		*/
	}
	return data, fmt.Errorf("Parse not implemented")
}

func retrieveResultPage(link string) (string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", body), nil
}

// Use a simple regexp to pull this data from the head element. The 
func getCitationPdfUrl(src string) (string, error) {
	re := regexp.MustCompile(`<meta[^>]+name="citation_pdf_url"[^>]+content="([^"]+)">`)
	matches := re.FindStringSubmatch(src)
	if len(matches) < 2 {
        return "", fmt.Errorf("missing meta[name='citation_pdf_url']")
    }

	return matches[1], nil
}

func retrievePatentPdf(link string) error {
	u, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("failed to parse citation pdf url %q, %s", link, err)
	}
	fName := path.Base(u.Path)
	if fName == "" {
		return fmt.Errorf("failed to determine PDF filename, %q", u.Path)
	}
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to ready body from %q, %s", link, err)
	}
	if err := os.WriteFile(fName, body, 0664); err != nil {
		return fmt.Errorf("failed to write %s, %s", fName, err)
	}
	return nil
}

func Process(data []*PatentRecord) error {
	errCnt := 0
	for i, rec := range data {
		if rec.ResultLink != "" {
			src, err := retrieveResultPage(rec.ResultLink)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to retrieve (%d) %q, %s\n", i+2, rec.ResultLink, err)
				errCnt++
				continue
			}
			citationPdfUrl, err := getCitationPdfUrl(src)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to find citation PDF URL (%+d), %s", i+2, err)
				errCnt++
				continue
			}
			if err := retrievePatentPdf(citationPdfUrl); err != nil {
				fmt.Fprintf(os.Stderr, "failed to retreive (%d) %q, %s\n", i+2, citationPdfUrl, err)
				continue
			}
		} else {
			fmt.Printf("skipping row %d, no result link value\n", i+2)
		}
	}
	if errCnt > 0 {
		return fmt.Errorf("%d errors processing CSV file")
	}
	return nil
}

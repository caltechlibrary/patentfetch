

# PatentFetch

PatentFetch uses a CSV file provided by Google's Patent search download link to retrieve patent PDFs if available.

- The individual search result page's HTML is retrieved and parsed for the meta element containing the `name="citation_pdf_url"`. 
- The citation pdf url is used to then retrieve the PDF of the patent

## Release Notes

- version: 0.0.1
- status: concept
- released: 2025-10-17

Proof of concept command line tool to retrieving patent PDFs for Caltech Archives.


### Authors

- Doiel, R. S.



## Software Requirements

- go >= 1.25.3

### Software Suggestions

- CMTools >= 0.0.40



## Related resources



- [Getting Help, Reporting bugs](https://github.com/caltechlibrary/PatentFetch/issues)
- [LICENSE](https://caltechlibrary.github.io/PatentFetch/LICENSE)
- [Installation](INSTALL.md)
- [About](about.md)


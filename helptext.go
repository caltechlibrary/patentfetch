package patentfetch

var (
	HelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] CSV_FILENAME

# DESCRIPTION

{app_name} will retrieve the provided CSV file, use the "result link" column to fetch the
the results landing page for the patent in the row and parse that page for the meta element
with the name value of "citation_pdf_url". If that is found then the PDF will be retrieved too.

On success the CSV file will be updated for the filename retreived, on failure that column value
will be set to an empty string.

# OPTIONS

-help
: display this help page

-version
: display version info

-license
: display license

# EXAMPLE

~~~shell
{app_name} gp-search-20251017-143309.csv
~~~

This will process the CSV file and retrieve the PDFs it can find.

`
)

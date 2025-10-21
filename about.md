---
title: PatentFetch
abstract: "PatentFetch uses a CSV file provided by Google&#x27;s Patent search download link to retrieve patent PDFs if available.

- The individual search result page&#x27;s HTML is retrieved and parsed for the meta element containing the &#x60;name&#x3D;&quot;citation_pdf_url&quot;&#x60;. 
- The citation pdf url is used to then retrieve the PDF of the patent"
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903



repository_code: https://github.com/caltechlibrary/PatentFetch.git
version: 0.0.1
license_url: https://caltechlibrary.github.io/PatentFetch/LICENSE

programming_language:
  - Go


date_released: 2025-10-17
---

About this software
===================

## PatentFetch 0.0.1

Proof of concept command line tool to retrieving patent PDFs for Caltech Archives.

### Authors

- R. S. Doiel, <https://orcid.org/0000-0003-0900-6903>






PatentFetch uses a CSV file provided by Google&#x27;s Patent search download link to retrieve patent PDFs if available.

- The individual search result page&#x27;s HTML is retrieved and parsed for the meta element containing the &#x60;name&#x3D;&quot;citation_pdf_url&quot;&#x60;. 
- The citation pdf url is used to then retrieve the PDF of the patent

- License: <https://caltechlibrary.github.io/PatentFetch/LICENSE>
- GitHub: <https://github.com/caltechlibrary/PatentFetch.git>
- Issues: <https://github.com/caltechlibrary/PatentFetch/issues>

### Programming languages

- Go




### Software Requirements

- go >= 1.25.3


### Software Suggestions

- CMTools &gt;&#x3D; 0.0.40



Installation for development of **PatentFetch**
===========================================

**PatentFetch** PatentFetch uses a CSV file provided by Google's Patent search download link to retrieve patent PDFs if available.

- The individual search result page's HTML is retrieved and parsed for the meta element containing the `name="citation_pdf_url"`. 
- The citation pdf url is used to then retrieve the PDF of the patent

Quick install with curl or irm
------------------------------

There is an experimental installer.sh script that can be run with the following command to install latest table release. This may work for macOS, Linux and if you’re using Windows with the Unix subsystem. This would be run from your shell (e.g. Terminal on macOS).

~~~shell
curl https://caltechlibrary.github.io/PatentFetch/installer.sh | sh
~~~

This will install the programs included in PatentFetch in your `$HOME/bin` directory.

If you are running Windows 10 or 11 use the Powershell command below.

~~~ps1
irm https://caltechlibrary.github.io/PatentFetch/installer.ps1 | iex
~~~

### If your are running macOS or Windows

You may get security warnings if you are using macOS or Windows. See the notes for the specific operating system you're using to fix issues.

- [INSTALL_NOTES_macOS.md](INSTALL_NOTES_macOS.md)
- [INSTALL_NOTES_Windows.md](INSTALL_NOTES_Windows.md)

Installing from source
----------------------

### Required software

- go &gt;&#x3D; 1.25.3

### Steps

1. git clone https://github.com/caltechlibrary/PatentFetch
2. Change directory into the `PatentFetch` directory
3. Make to build, test and install

~~~shell
git clone https://github.com/caltechlibrary/PatentFetch
cd PatentFetch
make
make test
make install
~~~


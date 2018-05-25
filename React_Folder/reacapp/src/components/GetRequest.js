export default function GetReportXML(sUrl) {
    var XMLTable = ""
    var xhttp = new XMLHttpRequest();
    
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            XMLTable = JSON.parse(this.responseText);
            console.log(XMLTable)
        }
    };
  
    xhttp.open("GET", "http://localhost:8080"+sUrl, false);
    xhttp.send(); 
  
    return XMLTable
  }
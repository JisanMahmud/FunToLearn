export default function VarifyLoginCred(sUrl) {
  var LoggedIn = false;
  var xhttp = new XMLHttpRequest();
  
  xhttp.onreadystatechange = function() {
      if (this.readyState === 4 && this.status === 200) {
        LoggedIn = true;
      }
  };

  xhttp.open("POST", "http://localhost:8080"+sUrl, false);
  xhttp.send(); 

  return LoggedIn

}
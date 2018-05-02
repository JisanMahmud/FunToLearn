import WebRequest from 'web-request'
// import rp from 'request'

export default async function VarifyLoginCred(sUrl) {
  
  var None = "";
  var LoggedIn = false;
  var result = await WebRequest.post("http://localhost:8080/"+sUrl, null, None);
  console.log("This is from web request: "+ result.content);
  console.log(result);
  if(result.content === 'LoggedIn\n') LoggedIn = true;
  console.log("this is after responce:" +LoggedIn);
  return LoggedIn

  // var options = {
  //   method: 'POST',
  //   uri: "http://localhost:8080/"+sUrl,
  //   headers: {
  //       /* 'content-type': 'application/x-www-form-urlencoded' */ // Is set automatically
  //   }
  // };

  // rp(options)
  //     .then(function (body) {
  //       console.log(body);
  //         // POST succeeded...
  //     })
  //     .catch(function (err) {
  //         // POST failed...
  //   });


}
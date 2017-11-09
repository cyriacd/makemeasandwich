// Require sudo on unix platforms
if(process.platform !== 'win32'){
  if(process.env.USER !== "root"){
    console.log("What? Make it yourself.");
    process.exit();
  }else{
    console.log("Sandwich made. Go get it!");
  }
}

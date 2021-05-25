function updatetime() {
  var c = document.getElementById("realtime");

  c.innerHTML =
    new Date().toLocaleDateString() + " " + new Date().toLocaleTimeString();
}
setInterval(updatetime, 500);

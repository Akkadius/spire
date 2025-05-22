/**
 * dark-mode.js
 */

if (typeof (Storage) !== "undefined") {
  loadDarkMode();
} else {
  // Sorry! No Web Storage support..
}

function toggleDarkMode() {
  var darkMode = localStorage.getItem("dark_mode");
  if (typeof darkMode !== "undefined" && darkMode === "1") {
    setLightMode();
    localStorage.setItem("dark_mode", "0");
    $(".night-mode-toggle").prop("checked", false);
    // console.log("Dark mode is now %s", localStorage.getItem("dark_mode"));
    return false;
  }

  setDarkMode();

  $(".night-mode-toggle").prop("checked", true);

  localStorage.setItem("dark_mode", "1");

  // console.log("Dark mode is now %s", localStorage.getItem("dark_mode"));
}

function loadDarkMode() {
  var darkMode = localStorage.getItem("dark_mode");
  if (typeof darkMode !== "undefined" && darkMode === "1") {
    setDarkMode();
    return false;
  }
  setLightMode();
}

function setDarkMode() {
  $("style[dark-mode]")
    .html("html { filter: invert(100%); } body { background-color: rgba(0, 0, 0, 0.97) !important; background-blend-mode: darken;} " +
      "img, .avatar { filter: invert(100%); }");
}

function setLightMode() {
  console.log("light mode");
  $("style[dark-mode]")
    .html("html { filter: none; } body { background-color: rgba(245, 247, 251, 0.99 ) !important; background-blend-mode: overlay; } " +
      "img, .avatar { filter: invert(0%); }");
}

function checkLightModeToggleInput() {
  var darkMode = localStorage.getItem("dark_mode");
  // console.log("dark mode is %s", darkMode);
  if (typeof darkMode !== "undefined" && darkMode === "1") {
    $(".night-mode-toggle").prop("checked", true);
    return false;
  }
}

window.addEventListener("load", function(event) {
  checkLightModeToggleInput();
  $(".custom-switch").fadeIn(200);
});

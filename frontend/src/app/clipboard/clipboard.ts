export default class Clipboard {
  static copyFromText(text) {
    let tempNode   = document.createElement("input");
    tempNode.type  = "text";
    tempNode.value = text;
    document.body.appendChild(tempNode);
    tempNode.select();
    document.execCommand("Copy");
    document.body.removeChild(tempNode);
  }

  static copyFromElement(element) {
    // @ts-ignore
    const str = document.getElementById(element).innerText
    const el  = document.createElement('textarea')
    el.value  = str
    el.setAttribute('readonly', '')
    el.style.position = 'absolute'
    el.style.left     = '-9999px'
    document.body.appendChild(el)
    el.select()
    document.execCommand('copy')
    document.body.removeChild(el)
  }
}

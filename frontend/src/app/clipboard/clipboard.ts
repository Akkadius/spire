export default class Clipboard {
  static copyFromText(text) {
    const el  = document.createElement('textarea')
    el.value  = text
    el.setAttribute('readonly', '')
    el.style.position = 'absolute'
    el.style.left     = '-9999px'
    document.body.appendChild(el)
    el.select()
    document.execCommand('copy')
    document.body.removeChild(el)
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

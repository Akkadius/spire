export class EditFormFieldUtil {

  static setFieldHighlightHasSubEditor(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-green')
    }
  }

  static resetFieldHighlightHasSubEditorStatus() {
    document.querySelectorAll("input, select, textarea").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-green')) {
        element.classList.remove('pulsate-highlight-green')
      }
    });
  }

  static resetFieldSubEditorHighlightedStatus() {
    document.querySelectorAll("input, select, textarea").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-white')) {
        element.classList.remove('pulsate-highlight-white')
      }
    });
  }

  static anyFieldsHaveBeenEdited() {
    let edited = false
    document.querySelectorAll("input, select, textarea").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-modified')) {
        edited = true
      }
    });
    return edited
  }

  static resetFieldEditedStatus() {
    document.querySelectorAll("input, select, textarea").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-modified')) {
        element.classList.remove('pulsate-highlight-modified')
      }
    });
  }

  static setFieldModified(evt) {
    if (parseInt(evt.target.getAttribute('ignore-input-change')) === 1) {
      return
    }

    evt.target.classList.add('pulsate-highlight-modified');
  }

  static setFieldModifiedById(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-modified')
    }
  }

  static clearFieldModifiedById(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.remove('pulsate-highlight-modified')
    }
  }

  static setFieldSubEditorHighlightedById(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-white')
    }
  }
}

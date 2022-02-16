export class EditFormFieldUtil {

  static setFieldHighlightHasSubEditor(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-green')
    }
  }

  static resetFieldHighlightHasSubEditorStatus() {
    document.querySelectorAll("input, select").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-green')) {
        element.classList.remove('pulsate-highlight-green')
      }
    });
  }

  static resetFieldSubEditorHighlightedStatus() {
    document.querySelectorAll("input, select").forEach((element) => {
      if (element && element.classList.contains('pulsate-highlight-white')) {
        element.classList.remove('pulsate-highlight-white')
      }
    });
  }

  static resetFieldEditedStatus() {
    document.querySelectorAll("input, select").forEach((element) => {
      // @ts-ignore
      element.style.setProperty('border-color', '#555555', 'important');
    });
  }

  static setFieldModified(evt) {
    evt.target.style.setProperty('border-color', 'orange', 'important');
  }

  static setFieldModifiedById(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-modified')
    }
  }

  static setFieldSubEditorHighlightedById(id) {
    const target = document.getElementById(id)
    if (target) {
      target.classList.add('pulsate-highlight-white')
    }
  }
}

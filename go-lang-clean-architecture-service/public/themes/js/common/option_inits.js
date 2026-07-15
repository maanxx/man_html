function initializeCodeMirror(textareaId) {
    return CodeMirror.fromTextArea(document.getElementById(textareaId), {
        lineNumbers: true,
        viewportMargin: Infinity,
        mode: "text/html",
        matchBrackets: true,
        lineWrapping: true,
        autoCloseTags: true,
        showAutoCompleteButton: true,
        foldGutter: true,
        gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"]
    });
}

export {
    initializeCodeMirror,
};
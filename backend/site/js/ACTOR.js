
const actorForm = document.actorForm;
const addButton = actorForm.addButton, 
    removeButton = actorForm.removeButton, 
    languagesSelect = actorForm.language;
// обработчик добавления элемента
function addOption(){
    const text = actorForm.textInput.value;
    const value = actorForm.valueInput.value;

    const newOption = new Option(text, value);
    languagesSelect.options[languagesSelect.options.length]=newOption;
}

function removeOption(){
    const selectedIndex = languagesSelect.options.selectedIndex;
    languagesSelect.options[selectedIndex] = null;
}

addButton.addEventListener("click", addOption);
removeButton.addEventListener("click", removeOption);


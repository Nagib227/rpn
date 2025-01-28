document.getElementById('search-input').addEventListener('keypress', function(event) {
    // Проверяем, была ли нажата клавиша Enter
    console.log(event.key)
    if (event.key === "Enter") {
        // Получаем значение из поля ввода
        const searchTerm = document.getElementById('search-input').value;
        const content = document.getElementById('widget');
        
        // Удаляем предыдущие выделения
        removeHighlights(content);
        
        if (searchTerm) {
            // Выполняем поиск и выделяем найденные слова
            highlightWords(content, searchTerm);
        }
    }
});

// Функция для выделения слов
function highlightWords(container, searchTerm) {
    const regex = new RegExp(`(${searchTerm})`, 'gi'); // Регулярное выражение для поиска
    container.innerHTML = container.innerHTML.replace(regex, '<span class="highlight">$1</span>');
}

// Функция для удаления выделений
function removeHighlights(container) {
    const highlighted = container.querySelectorAll('.highlight');
    highlighted.forEach(function(node) {
        const parent = node.parentNode;
        parent.replaceChild(document.createTextNode(node.innerText), node);
    });
}
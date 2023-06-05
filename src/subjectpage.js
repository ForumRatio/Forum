window.addEventListener('DOMContentLoaded', function() {
    var categories = document.querySelectorAll('.categories .subtitle');
    
    for (var i = 0; i < categories.length; i++) {
        if (i >= 4) {
            categories[i].style.display = 'none';
        }
    }
});

document.querySelector('.tobecontinued img.normal').addEventListener('click', function() {
    window.location.href = 'page-suivante.html';
});

document.querySelector('.tobecontinued_left img.left').addEventListener('click', function() {
    window.history.back();
});

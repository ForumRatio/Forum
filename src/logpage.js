let connexion = document.getElementsByClassName('connexion')
let inscription = document.getElementsByClassName('inscription')
connexion[0].addEventListener('click', () => {
    window.location="/login"
 });
inscription[0].addEventListener('click', () => {
    window.location="/register"
});
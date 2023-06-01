const triangle = document.getElementById('triangle');
const profileImage = document.getElementById('profileImage');
let imageChanged = false;
        
triangle.addEventListener('click', () => {
    if (imageChanged) {
        profileImage.src = '../asset/morganapdp.png';
        imageChanged = false;
    } else {
        profileImage.src = '../asset/igypdp.png';
        imageChanged = true;
    }
});
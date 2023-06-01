const triangle = document.getElementById('triangle');
const profileImage = document.getElementById('profileImage');
let image = ["morganapdp.png","igypdp.png"]
let imageChanged = 0;
        
triangle.addEventListener('click', () => {
    if (imageChanged+1 == image.length){
        imageChanged = 0
    } else {
        imageChanged++
    }
    profileImage.src = '../asset/' + image[imageChanged]
});
let button = document.getElementsByClassName('button')
let input = document.querySelectorAll('input')
function Create(){
    if (input[2] === input[3]) {
        fetch('/createUser',{
            method: "POST",
            headers: {"content-type":"application/json"},
            body: JSON.stringify({
            Username : input[0].value,
            Phone: input[1].value,
            Password : input[2].value,
            Mail: input[4].value,
            }) 
        })

        .then((re) => {
            console.log(re)
            if (re.redirected){
                window.location="/categorypage"
            }
        })
    }
}
button[0].addEventListener('click', () => {
    Create();
    // window.location="/categorypage"
 });
console.log(button[0])
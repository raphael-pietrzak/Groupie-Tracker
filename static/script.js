

console.log(art)

function search() {
    let art = document.getElementsByClassName("card")
    let search_input = document.getElementById("searchBar").value;
    for (i = 0; i < art.length; i++) {
        console.log("hi")
        if  (art[i].innerHTML.toLowerCase().includes(search_input)) {
            art[i].style.display = 'block';
        } else {
            art[i].style.display = 'none';
        }
    }
}
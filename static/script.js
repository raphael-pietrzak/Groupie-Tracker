

function search() {
    const noResult = "No Results"
    let art = document.getElementsByClassName("test")
    let search_input = document.getElementById("searchBar").value;
    for (i = 0; i < art.length; i++) {
        console.log(art[i].innerHTML)
        if  (art[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
            art[i].style.display = 'block';
        } else {
            art[i].style.display = 'none';
        }
    }
}
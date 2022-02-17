function search() {
    let art = document.getElementsByClassName("button");
    let searchBar = document.getElementsByClassName("nav");
    let search_input = document.getElementById("searchBar").value;
    let name = document.getElementsByName("artistName");
    let members = document.getElementsByName("members");
    let date = document.getElementsByName("date");
    let nothing = true
    document.getElementById("answer").innerHTML = ""

    for (i = 0; i < art.length; i++) {
      let j = i+1
      if (name[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          nothing = false;
          document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w="+ j +">" + name[i].innerHTML + "  name"+ "</a> <br>"
      } else if (members[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          nothing = false;
          document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w="+ j +">" + members[i].innerHTML + "  members"+ "</a> <br>"
      } else if (date[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w="+ j +">" + date[i].innerHTML + "  date"+ "</a> <br>"
          nothing = false;
      } else {
          art[i].style.display = 'none';
          
      }
    }
    if (search_input == "") {
      document.getElementById("answer").innerHTML = ""
      searchBar[0].style.borderRadius = '70px';
    } else {
      searchBar[0].style.borderRadius = '8px';
    }
    if (nothing) {
      document.getElementById("noresult").innerHTML = "We couldn't find any matches for  \" " + search_input + " \"";
    } else {
      document.getElementById("noresult").innerHTML = "";
    }
}

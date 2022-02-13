function search() {
    let art = document.getElementsByClassName("test");
    let search_input = document.getElementById("searchBar").value;
    let name = document.getElementsByName("s1");
    let members = document.getElementsByName("members");
    let date = document.getElementsByName("date");
    let nothing = true
    document.getElementById("answer").innerHTML = ""
    for (i = 0; i < art.length; i++) {
      if (name[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          nothing = false;
          document.getElementById("answer").innerHTML += "<li>" + name[i].innerHTML + "  name"+ "</li>"
      } else if (members[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          nothing = false;
          document.getElementById("answer").innerHTML += "<li>" + members[i].innerHTML + "  members"+ "</li>"
      } else if (date[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
          document.getElementById("answer").innerHTML += "<li>" + date[i].innerHTML + "  date"+ "</li>"
          nothing = false;
      } else{
          art[i].style.display = 'none';
          
      }
    }
    if (nothing) {
      document.getElementById("noresult").innerHTML = "We couldn't find any matches for  \" " + search_input + " \"";
    } else {
      document.getElementById("noresult").innerHTML = "";
    }
}

function numb_members(num) {
  if (document.getElementById(num).checked == true) {
    console.log(num)
  }
}

function dateCreation_filter() {
  let creation = document.getElementById("dateCreation").value;
  console.log(creation)
}

function firstAlb_filter() {
  let album = document.getElementById("FirstAlb").value;
  console.log(album)
}

function searchFilter() {
    document.getElementById("myDropdown").classList.toggle("show");
  }
  
  // Close the dropdown if the user clicks outside of it
  window.onclick = function(event) {
    if (!event.target.matches('.dropbtn')) {
      var dropdowns = document.getElementsByClassName("dropdown-content");
      var i;
      for (i = 0; i < dropdowns.length; i++) {
        var openDropdown = dropdowns[i];
        // if (openDropdown.classList.contains('show')) {
        //   // openDropdown.classList.remove('show');
        // }
      }
    }
  }
function search() {
    let art = document.getElementsByClassName("test");
    let search_input = document.getElementById("searchBar").value;
    for (i = 0; i < art.length; i++) {
      let minusMembers = art[i].innerHTML.toLowerCase().replace('','');
      minusMembers = minusMembers.replace('members','');
      console.log(minusMembers)
      if (minusMembers.toLowerCase().includes(search_input.toLowerCase())) {
          art[i].style.display = 'block';
      } else {
          art[i].style.display = 'none';
      }
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
var slider = document.getElementById("FirstAlb");
var output = document.getElementById("demo");
output.innerHTML = slider.value;

slider.oninput = function() {
  output.innerHTML = this.value;
}

function search() {
  document.getElementsByClassName("dropdown-content")[0].style.display = 'none';
  let searchBar = document.getElementsByClassName("nav");
  let search_input = document.getElementById("searchBar").value;

  let art = document.getElementsByClassName("button");

  let name = document.getElementsByName("artistName");

  let members2 = [['Freddie Mercury','Brian May','John Daecon','Roger Meddows-Taylor','Mike Grose','Barry Mitchell','Doug Fogie'],['Jacob Hemphill','Bob Jefferson','Ryan "Byrd" Berty','Ken Brownell','Patrick O\'Shea','Hellman Escorcia','Rafael Rodriguez','Trevor Young'],['David Gilmour','Syd Barrett','Roger Waters','Richard Wright','Nick Manson','Bob Klose'],['Klaus Meine','Rudolf Schenker','Matthias Jabs','Pawel Maciwoda','Mikkey Dee'],['Jahseh Dwayne Ricardo Onfroy'],['Malcolm James McCormick'],['Gary Maurice Lucas Jr.'],['Kendrick Lamar Duckworth'],['Angus Young','Chris Slade','Stevie Young','Axl Rose'],['Eddie Vedder','Mike McCready','Stone Gossard','Jeff Ament','Matt Cameron'],['Katheryn Elizabeth Hudson'],['Robyn Rihanna Fenty'],['Mike Rutherford','Tony Banks','Phil Collins','Chester Thompson','Daryl Stuermer'],['Phil Collins'],['Jimmy Page','John Paul Jones','Robert Plant','John Bonham'],['Jimi Hendrix','Mitch Mitchell','Noel Redding','Billy Cox'],['Barry Gib','Robin Gibb','Maurice Gibb'],['Ian Paice','Ian Gillan','Roger Glover','Steve Morse','Don Airey'],['Steven Tyler','Tom Hamilton','Joey Kramer','Joe Perry','Brad Whitford'],['Mark Knopfler','David Knopfler','John Illsley','Pick Withers'],['Dinho','Bento Hinoto','Júlio Rasec','Sérgio Reoli','Samuel Reis de Oliveira'],['Jared Leto','Shannon Leto','Tomo Miličević','Matt Wachter','Solon Bixler'],['Dan Reynolds','Daniel Wayne Sermon','Ben McKee'],[' Jarad Higgins'],['Sir Robert Bryson Hall II'],['Alec Benjamin'],['Bobby McFerrin'],['Fadil El Ghoul'],['Post Malone'],['Jacques Bermon Webster II'],['Jermaine Lamarr Cole'],['Chad Kroeger','Ryan Peake','Mike Kroeger','Daniel Adair','Brandon Kroeger','Ryan Vikedal','Mitch Guindon'],['Prodigy','Havoc'],['Axl Rose','Duff McKagan','Slash','Dizzy Reed','Richard Fortus','Frank Ferrer','Melissa Reese'],['Ice Cube','Eazy-E','Dr. Dre','MC Ren','DJ Yella','The D.O.C.','Arabian Prince'],['Bono','The Edge','Adam Clayton','Larry Mullen Jr.','Ivan McCormick','Peter Martin','Dick Evans'],['Alex Turner','Matt Helders','Jamie Cook','Nick O\'Malley','Andy Nicholson','Glyn Jones'],['Pete Wentz','Patrick Stump','Andy Hurley','Joe Trohman'],['Damon Albarn','Jamie Hewlett','Remi Kabaka Jr.'],['Glenn Frey','Don Henley','Randy Meisner','Bernie Leadon'],['Chester Bennington','Mike Shinoda','Brad Delson','Dave Farrell','Joe Hahn','Rob Bourdon'],['Anthony Kiedis','Flea','Chad Smith','Josh Klinghoffer'],['Marshall Bruce Mathers'],['Billie Joe Armstrong','Mike Dirnt','Tré Cool'],['James Hetfield','Lars Ulrich','Kirk Hammett','Robert Trujillo'],['Chris Martin','Guy Berryman','Jonny Buckland','Will Champion'],['Adam Levine','Jesse Carmichael','Mickey Madden','James Valentine','Matt Flynn','PJ Morton','Sam Farrar'],['Tyler Joseph','Josh Dun'],['Mick Jagger ','Keith Richards','Charlie Watts','Ron Wood'],['Matthew Bellamy','Christopher Wolstenholme','Dominic Howard '],['Dave Grohl','Nate Mendel','Taylor Hawkins','Chris Shiflett','Pat Smear','Rami Jaffee'],['Alexander Pall','Andrew Taggart','Matt McGuire','Tony Ann']]
  let date = document.getElementsByName("date");
  let nothing = true

  document.getElementById("answer").innerHTML = ""
  for (i = 0; i < art.length; i++) {

    // for (k = 0; k < members[i].length; k++) {
    //   if (members[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
    //     mmm = members[i];
    //     console.log(mmm.innerHTML)
    //   }
    // }
    let a = ""
    for (k = 0; k < members2[i].length; k++){
      if (members2[i][k].toLowerCase().includes(search_input.toLowerCase())){
        a = members2[i][k]
        break
      }
    }

    let j = i + 1;
    if (name[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
      art[i].style.display = 'block';
      nothing = false;
      document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w=" + j + ">" + name[i].innerHTML + "<p class=" + "typesearch" + ">  name</p>" + "</a> ";
    } else if (a != "") {
      art[i].style.display = 'block';
      nothing = false;
      document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w=" + j + ">" + a + "<p class=" + "typesearch" + ">  members</p>" + "</a>";
    } else if (date[i].innerHTML.toLowerCase().includes(search_input.toLowerCase())) {
      art[i].style.display = 'block';
      document.getElementById("answer").innerHTML += "<a href=http://localhost:5500/artist?w=" + j + ">" + date[i].innerHTML + "<p class=" + "typesearch" + ">  date</p>" + "</a> ";
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
  if (document.getElementById("answer").innerHTML == "") {
    searchBar[0].style.borderRadius = '70px';
  }
  if (nothing) {
    document.getElementById("noresult").innerHTML = "We couldn't find any matches for  \"" + search_input + "\"";
  } else {
    document.getElementById("noresult").innerHTML = "";
  }
}

function searchFilter() {

  document.getElementById("myDropdown").classList.toggle("show");
  document.getElementById("searchBar").value = "";
  document.getElementById("answer").innerHTML = ""
  let searchBar = document.getElementsByClassName("nav");
  searchBar[0].style.borderRadius = '70px';
  document.getElementsByClassName("dropdown-content")[0].style.display = 'block';
}


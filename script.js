// Search Bar : Artists/band name, members, locations, first album date, creation date

const charactersList = document.getElementById('charactersList');
const searchBar = document.getElementById('searchBar');
let hpCharacters = [];

searchBar.addEventListener('keyup', (e) => {
    const searchString = e.target.value.toLowerCase();

    const filteredCharacters = hpCharacters.filter((character) => {
        return (
            character.name.toLowerCase().includes(searchString) || character.members.toLowerCase().includes(searchString)
        );
    });
    displayCharacters(filteredCharacters);
});

const displayCharacters = (characters) => {
    const htmlString = characters
        .map((character) => {
            return `
            <li class="character">
                <h2>${character.Name}</h2>
                <p>House: ${character.members}</p>
                <img src="${character.image}"></img>
            </li>
        `;
        })
        .join('');
    charactersList.innerHTML = htmlString;
};
loadCharacters();


// const artists = []

// const people = [
//     { name: 'Queen'},
//     { name: 'SOJA'},
//     { name: 'Pink Floyd'},
//     { name: 'Scorpions'},
//     { name: 'XXXTentacion'},
//     { name: 'Mac Miller'},
//     { name: 'Joyner Lucas'},
//     { name: 'Kendrick Lamar'},
//     { name: 'ACDC'},
//     { name: 'Pearl Jam'},
//     { name: 'Katy Perry'},
//     { name: 'Rihanna'},
//     { name: 'Genesis'},
//     { name: 'Phil Collins'},
//     { name: 'Led Zeppelin'},
//     { name: 'The Jimi Hendrix Experience'},
//     { name: 'Bee Gees'},
//     { name: 'Deep Purple'},
//     { name: 'Aerosmith'},
//     { name: 'Dire Straits'},
//     { name: 'Mamonas Assassinas'},
//     { name: 'Thirty Seconds to Mars'},
//     { name: 'Imagine Dragons'},
//     { name: 'Juice Wrld'},
//     { name: 'Logic'},
//     { name: 'Alec Benjamin'},
// ]

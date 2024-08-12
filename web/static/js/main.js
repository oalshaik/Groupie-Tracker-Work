document.addEventListener("DOMContentLoaded", function() {
    const loading = document.getElementById('loading');
    const artistsList = document.getElementById('artists-list');
    const searchInput = document.getElementById('search');

    let allArtists = [];  // Store all artists for filtering

    // Show loading indicator
    loading.style.display = 'block';

    fetch('/artists')
        .then(response => response.json())
        .then(artists => {
            allArtists = artists;  // Store all artists
            renderArtists(artists);
            loading.style.display = 'none';
            artistsList.style.display = 'grid';
        })
        .catch(error => {
            // Show error message instead of artists list
            loading.textContent = 'Failed to load artists. Please try again later.';
            console.error('Error fetching artists:', error);
        });

    // Function to render artists
    function renderArtists(artists) {
        artistsList.innerHTML = '';  // Clear the list
        artists.forEach(artist => {
            const artistDiv = document.createElement('div');
            artistDiv.className = 'artist';

            const artistImage = document.createElement('img');
            artistImage.src = artist.image;  // Use correct case (image)
            artistDiv.appendChild(artistImage);

            const artistName = document.createElement('div');
            artistName.className = 'artist-name';
            artistName.textContent = artist.name;  // Use correct case (name)
            artistDiv.appendChild(artistName);

            const artistAlbum = document.createElement('div');
            artistAlbum.className = 'artist-album';
            artistAlbum.textContent = `First Album: ${artist.firstAlbum}`;  // Use correct case (firstAlbum)
            artistDiv.appendChild(artistAlbum);

            artistDiv.addEventListener('click', () => {
                window.location.href = `/artists/${artist.id}`;  // Use correct case (id)
            });

            artistsList.appendChild(artistDiv);
        });
    }

    // Event listener for search input
    searchInput.addEventListener('input', function() {
        const filteredArtists = allArtists.filter(artist => 
            artist.name.toLowerCase().includes(searchInput.value.toLowerCase())  // Use correct case (name)
        );
        renderArtists(filteredArtists);
    });
});

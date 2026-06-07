<div class="hero-canvas">
    <div class="hero-body">
        <h1>Discover your next destination</h1>
        <p>Search countries, explore attractions, and curate your personal travel wishlist.</p>
        
        <div class="search-module">
            <label for="country-search-input">WHERE TO NEXT?</label>
            <div class="input-wrapper">
                <input type="text" id="country-search-input" placeholder="e.g. Bangladesh" autocomplete="off">
                <div id="autocomplete-results-box" class="autocomplete-dropdown hidden"></div>
            </div>
        </div>
    </div>
</div>

<section class="section-container">
    <h2>Featured destinations</h2>
    <div class="destination-grid">
        {{range .FeaturedCountries}}
        <div class="destination-card">
            <div class="flag-frame">
                <img src="{{.Flags.Png}}" alt="{{.Name.Common}} Flag">
            </div>
            <div class="card-details">
                <h3>{{.Name.Common}}</h3>
                <p>{{range .Capital}}{{.}}{{end}} , {{.Region}}</p>
            </div>
        </div>
        {{end}}
    </div>
</section>

<section class="section-container">
    <h2>Popular attractions</h2>

</section>
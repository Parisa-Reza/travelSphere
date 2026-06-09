<div class="hero-canvas">
    <div class="hero-body">
        <h1>Discover your next destination</h1>
        <p>Search countries, explore attractions, and curate your personal travel wishlist.</p>
        
        <div class="search-module">
            <label for="country-search-input">WHERE TO NEXT?</label>
            <div class="input-wrapper">
                <input type="text" id="country-search-input" placeholder="e.g. dhaka or Argentina" autocomplete="off">
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

<div style="margin-top: 48px; font-family: system-ui, -apple-system, sans-serif; padding: 0 4px;">
    <h2 style="font-size: 22px; font-weight: 700; color: #0f172a; margin: 0 0 20px 0; letter-spacing: -0.3px;">Popular attractions</h2>
    
    <div style="display: flex; flex-direction: column; gap: 12px; max-width: 100%;">
        {{range .PopularAttractions}}
            <div style="background: #ffffff; border: 1px solid #e2e8f0; border-radius: 12px; padding: 18px 24px; display: flex; align-items: center; justify-content: space-between; box-shadow: 0 1px 2px rgba(0,0,0,0.02); transition: transform 0.15s ease, box-shadow 0.15s ease;">
                <div style="display: flex; align-items: baseline; gap: 8px;">
                    <span style="font-size: 16px; font-weight: 600; color: #1e293b;">{{.Name}}</span>
                    
                    {{if .DisplayKinds}}
                        <span style="font-size: 12px; color: #94a3b8; font-weight: 400; font-style: normal; margin-left: 4px;">
                            {{.DisplayKinds}}
                        </span>
                    {{end}}
                </div>
                
                <div style="color: #cbd5e1; font-size: 18px; font-weight: 400; padding-right: 4px; user-select: none;">
                    &#8250;
                </div>
            </div>
        {{else}}
            <div style="background: #f8fafc; border: 1px dashed #cbd5e1; border-radius: 12px; padding: 32px; text-align: center; color: #64748b; font-size: 14px;">
                Unable to display live local attractions at this time. Please check your connection or try again later.
            </div>
        {{end}}
    </div>
</div>
</section>
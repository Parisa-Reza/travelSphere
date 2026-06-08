<div style="padding: 40px; font-family: system-ui, -apple-system, sans-serif; background-color: #f8fafc; min-height: 100vh;">

    <div style="background: #ffffff; border-radius: 20px; border: 1px solid #e2e8f0; overflow: hidden; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.02); display: flex; gap: 40px; padding: 40px; margin-bottom: 32px;">
        <div style="width: 280px; flex-shrink: 0;">
            <img src="{{.Country.Flags.Png}}" alt="National Flag" style="width: 100%; border-radius: 12px; border: 1px solid #e2e8f0; display: block; object-fit: contain;">
        </div>
        
        <div style="flex-grow: 1;">
            <div style="display: inline-block; background-color: #e0e7ff; color: #4338ca; font-size: 11px; font-weight: 700; padding: 4px 12px; border-radius: 20px; text-transform: uppercase; letter-spacing: 0.5px; margin-bottom: 12px;">
                {{if .Country.Region}}{{.Country.Region}}{{else}}GLOBAL{{end}}
            </div>
            <h1 style="font-size: 38px; font-weight: 800; color: #0f172a; margin: 0 0 4px 0;">{{.Country.Name.Common}}</h1>
            <p style="color: #64748b; margin: 0 0 28px 0; font-size: 15px;">{{.Country.Name.Official}}</p>
            
            <div style="display: grid; grid-template-columns: repeat(4, 1fr); gap: 24px; border-top: 1px solid #f1f5f9; padding-top: 24px;">
                <div>
                    <h4 style="margin: 0; color: #94a3b8; font-size: 11px; font-weight: 700; letter-spacing: 0.5px; text-transform: uppercase;">CAPITAL</h4>
                    <p style="margin: 6px 0 0 0; color: #1e293b; font-size: 16px; font-weight: 600;">{{.Country.DisplayCapital}}</p>
                </div>
                <div>
                    <h4 style="margin: 0; color: #94a3b8; font-size: 11px; font-weight: 700; letter-spacing: 0.5px; text-transform: uppercase;">POPULATION</h4>
                    <p style="margin: 6px 0 0 0; color: #1e293b; font-size: 16px; font-weight: 600;">{{.Country.Population}}</p>
                </div>
                <div>
                    <h4 style="margin: 0; color: #94a3b8; font-size: 11px; font-weight: 700; letter-spacing: 0.5px; text-transform: uppercase;">CURRENCY</h4>
                    <p style="margin: 6px 0 0 0; color: #1e293b; font-size: 16px; font-weight: 600;">{{.Country.DisplayCurrencies}}</p>
                </div>
                <div>
                    <h4 style="margin: 0; color: #94a3b8; font-size: 11px; font-weight: 700; letter-spacing: 0.5px; text-transform: uppercase;">LANGUAGES</h4>
                    <p style="margin: 6px 0 0 0; color: #1e293b; font-size: 16px; font-weight: 600;">{{.Country.DisplayLanguages}}</p>
                </div>
            </div>
        </div>
    </div>

    <div style="background: #ffffff; border: 1px solid #e2e8f0; padding: 20px; border-radius: 12px; display: flex; align-items: center; justify-content: space-between;">
        <button id="add-wishlist-btn" style="background-color: #e11d48; color: #ffffff; border: none; padding: 12px 24px; font-weight: 600; border-radius: 8px; cursor: pointer; font-size: 14px;">
            Add to Wishlist
        </button>
        <div id="wishlist-feedback" style="font-size: 14px; font-weight: 500; color: #64748b;"></div>
    </div>
</div>
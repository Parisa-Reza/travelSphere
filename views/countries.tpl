<div style="padding: 24px; font-family: system-ui, -apple-system, sans-serif; background-color: #f8fafc; min-height: 100vh;">
    <div style="margin-bottom: 32px;">
        <h1 style="font-size: 32px; font-weight: 700; color: #0f172a; margin: 0 0 8px 0;">Country Explorer</h1>
    </div>

    <div style="display: flex; gap: 16px; background: #ffffff; padding: 24px; border-radius: 12px; border: 1px solid #e2e8f0; margin-bottom: 32px; box-shadow: 0 1px 2px rgba(0,0,0,0.05);">
        <div>
            <label style="display: block; font-size: 11px; font-weight: 700; color: #94a3b8; margin-bottom: 8px; letter-spacing: 0.5px;">SEARCH</label>
            <input type="text" id="explorer-search" placeholder="Country or capital..." style="padding: 12px 16px; border: 1px solid #cbd5e1; border-radius: 8px; width: 240px; background: #f8fafc; font-size: 14px; outline: none; color: #1e293b;">
        </div>
        <div>
            <label style="display: block; font-size: 11px; font-weight: 700; color: #94a3b8; margin-bottom: 8px; letter-spacing: 0.5px;">REGION</label>
            <select id="explorer-region-filter" style="padding: 12px 16px; border: 1px solid #cbd5e1; border-radius: 8px; width: 240px; background: #f8fafc; font-size: 14px; outline: none; color: #1e293b; cursor: pointer;">
                <option value="">All regions</option>
                <option value="Africa">Africa</option>
                <option value="Americas">Americas</option>
                <option value="Asia">Asia</option>
                <option value="Europe">Europe</option>
                <option value="Oceania">Oceania</option>
            </select>
        </div>
    </div>

    <div id="country-results" style="display: grid; grid-template-columns: repeat(auto-fill, minmax(260px, 1fr)); gap: 24px;">
        {{if .Error}}
            <div style="grid-column: 1/-1; color: #ef4444; font-weight: 500;">{{.Error}}</div>
        {{else}}
            {{range .Countries}}
            <a href="/countries/{{.Slug}}" style="text-decoration: none; color: inherit; display: block;">
                <div style="background: #ffffff; border: 1px solid #e2e8f0; border-radius: 16px; overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.02); height: 100%; display: flex; flex-direction: column;">
                    <div style="height: 160px; width: 100%; background: #f1f5f9;">
                        <img src="{{.Flags.Png}}" alt="Flag" style="width: 100%; height: 100%; object-fit: cover;" loading="lazy">
                    </div>
                    <div style="padding: 24px; flex-grow: 1;">
                        <h3 style="margin: 0 0 16px 0; font-size: 20px; font-weight: 700; color: #0f172a;">{{.Name.Common}}</h3>
                        <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Capital:</strong> {{.DisplayCapital}}</p>
                        <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Currency:</strong> {{.DisplayCurrencies}}</p>
                        <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Languages:</strong> {{.DisplayLanguages}}</p>
                    </div>
                </div>
            </a>
            {{else}}
                <div style="grid-column: 1/-1; color: #64748b; text-align: center; padding: 48px 0; font-size: 16px;">No countries matched the active criteria.</div>
            {{end}}
        {{end}}
    </div>
</div>

<script src="/static/js/countries.js"></script>
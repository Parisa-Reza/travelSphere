<div style="max-width: 1200px; margin: 40px auto; padding: 0 24px; font-family: system-ui, -apple-system, sans-serif;">
    <h1 style="font-size: 32px; font-weight: 800; color: #0f172a; margin: 0 0 8px 0;">Travel Wishlist</h1>
    <p style="color: #64748b; font-size: 15px; margin: 0 0 32px 0;">Edit notes, update trip status, or remove destinations. Changes save without reloading the page.</p>

    <div style="background: #ffffff; border: 1px solid #e2e8f0; border-radius: 16px; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.02); overflow: hidden;">
        <table style="width: 100%; border-collapse: collapse; text-align: left;">
            <thead>
                <tr style="background: #f8fafc; border-bottom: 1px solid #e2e8f0;">
                    <th style="padding: 16px 24px; font-size: 12px; font-weight: 600; color: #94a3b8; text-transform: uppercase;">Country</th>
                    <th style="padding: 16px 24px; font-size: 12px; font-weight: 600; color: #94a3b8; text-transform: uppercase;">Note</th>
                    <th style="padding: 16px 24px; font-size: 12px; font-weight: 600; color: #94a3b8; text-transform: uppercase;">Status</th>
                    <th style="padding: 16px 24px; font-size: 12px; font-weight: 600; color: #94a3b8; text-transform: uppercase;">Actions</th>
                </tr>
            </thead>
            <tbody>
                {{range .WishlistItems}}
                <tr id="row-{{.ID}}" style="border-bottom: 1px solid #e2e8f0;">
                    <td style="padding: 20px 24px; font-size: 16px; font-weight: 600; color: #1e293b;">
                        <a href="/countries/{{.Slug}}" style="color: inherit; text-decoration: none;">{{.CountryName}}</a>
                    </td>
                    <td style="padding: 20px 24px; width: 35%;">
                        <input type="text" id="note-{{.ID}}" value="{{.Note}}" placeholder="Add notes..." 
                               style="width: 100%; padding: 10px 14px; border: 1px solid #cbd5e1; border-radius: 8px; font-size: 14px; outline: none; background: #f8fafc;">
                    </td>
                    <td style="padding: 20px 24px;">
                        <select id="status-{{.ID}}" style="padding: 10px 14px; border: 1px solid #cbd5e1; border-radius: 8px; font-size: 14px; outline: none; background: #ffffff; cursor: pointer; min-width: 140px; font-weight: 500;">
                            <option value="Planned" {{if eq .Status "Planned"}}selected{{end}}>Planned</option>
                            <option value="Visited" {{if eq .Status "Visited"}}selected{{end}}>Visited</option>
                        </select>
                    </td>
                    <td style="padding: 20px 24px; white-space: nowrap;">
                        <button onclick="saveItem('{{.ID}}')" style="background: #e11d48; color: #ffffff; border: none; padding: 10px 20px; font-weight: 600; border-radius: 8px; cursor: pointer; font-size: 14px; margin-right: 8px; box-shadow: 0 4px 6px -1px rgba(225, 29, 72, 0.2);">
                            Save
                        </button>
                        <button onclick="deleteItem('{{.ID}}')" style="background: #ffffff; color: #e11d48; border: 1px solid #fca5a5; padding: 10px 16px; font-weight: 500; border-radius: 8px; cursor: pointer; font-size: 14px;">
                            Delete
                        </button>
                    </td>
                </tr>
                {{else}}
                <tr>
                    <td colspan="4" style="padding: 48px; text-align: center; color: #94a3b8; font-size: 15px;">
                        Your travel wishlist is currently empty.
                    </td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
</div>

<script src="/static/js/wishlist.js"></script>
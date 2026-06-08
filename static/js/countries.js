document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("explorer-search");
    const regionFilter = document.getElementById("explorer-region-filter");
    const resultsContainer = document.getElementById("country-results");

    async function fetchUpdatedGrid() {
        resultsContainer.innerHTML = `<div style="grid-column: 1/-1; color: #64748b; text-align: center; padding: 48px; font-size: 16px;">Loading matched travel targets...</div>`;

        const search = encodeURIComponent(searchInput.value.trim());
        const region = encodeURIComponent(regionFilter.value);

        try {
            const response = await fetch(`/api/countries?search=${search}&region=${region}`);
            if (!response.ok) throw new Error("API State Error");

            const list = await response.json();

            if (!list || list.length === 0) {
                resultsContainer.innerHTML = `<div style="grid-column: 1/-1; color: #64748b; text-align: center; padding: 48px; font-size: 16px;">No countries matched the active criteria.</div>`;
                return;
            }

            resultsContainer.innerHTML = list.map(c => `
                <a href="/countries/${c.Slug}" style="text-decoration: none; color: inherit; display: block;">
                    <div style="background: #ffffff; border: 1px solid #e2e8f0; border-radius: 16px; overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.02); height: 100%; display: flex; flex-direction: column;">
                        <div style="height: 160px; width: 100%; background: #f1f5f9;">
                            <img src="${c.flags.png}" alt="Flag" style="width: 100%; height: 100%; object-fit: cover;" loading="lazy">
                        </div>
                        <div style="padding: 24px; flex-grow: 1;">
                            <h3 style="margin: 0 0 16px 0; font-size: 20px; font-weight: 700; color: #0f172a;">${c.name.common}</h3>
                            <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Capital:</strong> ${c.DisplayCapital}</p>
                            <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Currency:</strong> ${c.DisplayCurrencies}</p>
                            <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Population:</strong> ${c.Population}</p>
                            <p style="font-size: 14px; margin: 6px 0; color: #475569;"><strong style="color: #64748b; font-weight: 500; display: inline-block; width: 85px;">Languages:</strong> ${c.DisplayLanguages}</p>
                        </div>
                    </div>
                </a>
            `).join("");

        } catch (err) {
            resultsContainer.innerHTML = `<div style="grid-column: 1/-1; color: #ef4444; font-weight: 500; text-align: center; padding: 48px;">Failed to refresh filtered search data.</div>`;
        }
    }

    let inputDebounceTimer;
    searchInput.addEventListener("input", () => {
        clearTimeout(inputDebounceTimer);
        inputDebounceTimer = setTimeout(fetchUpdatedGrid, 220);
    });

    regionFilter.addEventListener("change", fetchUpdatedGrid);
});
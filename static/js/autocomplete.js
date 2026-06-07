document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.getElementById("country-search-input");
    const resultsBox = document.getElementById("autocomplete-results-box");
    const displayContainer = document.querySelector(".destination-grid"); 

    searchInput.addEventListener("input", async () => {
        const query = searchInput.value.trim();

        // hide the box if input is empty
        if (query.length === 0) {
            resultsBox.innerHTML = "";
            resultsBox.classList.add("hidden");
            return;
        }

        try {
            const response = await fetch(`/api/countries?search=${encodeURIComponent(query)}`);
            if (!response.ok) throw new Error("Network status validation error");
            
            const data = await response.json();
            
            if (data && data.length > 0) {
                resultsBox.innerHTML = data.map(item => `
                    <div class="autocomplete-item" data-slug="${item.slug}" data-label="${item.label}">
                        ${item.label}
                    </div>
                `).join("");
                resultsBox.classList.remove("hidden");
            } else {
                resultsBox.innerHTML = `<div class="autocomplete-item" style="color: #94a3b8;">No results matched</div>`;
                resultsBox.classList.remove("hidden");
            }
        } catch (err) {
            console.error("AJAX fetch failed:", err);
        }
    });

    // Close autocomplete lists if outside element is clicked
    document.addEventListener("click", (e) => {
        if (e.target !== searchInput && e.target !== resultsBox) {
            resultsBox.classList.add("hidden");
        }
    });

    // Handle click on autocomplete items
    resultsBox.addEventListener("click", async (e) => {
        const targetItem = e.target.closest(".autocomplete-item");
        if (targetItem && targetItem.dataset.slug) {
            const selectedLabel = targetItem.dataset.label;
            
            // Update the search input with the selected label (without reloading the page)
            searchInput.value = selectedLabel.split("—")[0].trim();
            
            // Hide the dropdown box instantly
            resultsBox.classList.add("hidden");

        }
    });
});
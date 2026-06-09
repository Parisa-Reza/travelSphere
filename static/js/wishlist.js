document.addEventListener("DOMContentLoaded", function () {
    // Country Page when user clicks on add to wishlist button
    const wishlistBtn = document.getElementById("add-wishlist-btn");
    if (wishlistBtn) {
        wishlistBtn.addEventListener("click", function () {
            const countryName = document.getElementById("country-name").innerText.trim();
            const feedback = document.getElementById("wishlist-feedback");

            feedback.style.color = "#64748b";
            feedback.innerText = "adding to the wishlist...";

            fetch('/api/wishlist', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name: countryName })
            })
                .then(async response => {
                    const data = await response.json();
                    if (!response.ok) throw new Error(data.error || "Failed to process selection");

                    feedback.style.color = "#16a34a";
                    feedback.innerText = "adding to the wishlist...";

                    setTimeout(() => {
                        window.location.href = "/wishlist";
                    }, 800);
                })
                .catch(error => {
                    feedback.style.color = "#dc2626";
                    feedback.innerText = error.message;
                });
        });
    }

    //  save and delete feature in Wishlist page 
    const wishlistTableBody = document.getElementById("wishlist-table-body");
    if (wishlistTableBody) {
        wishlistTableBody.addEventListener("click", function (event) {
            const target = event.target;

            // Check if the clicked element is a Action Button
            if (target.tagName === "BUTTON" && target.hasAttribute("data-id")) {
                const id = target.getAttribute("data-id");
                const action = target.getAttribute("data-action");

                if (action === "save") {
                    saveItem(id);
                } else if (action === "delete") {
                    deleteItem(id);
                }
            }
        });
    }
});

// Update  in the wishlist
function saveItem(id) {
    const noteElement = document.getElementById(`note-${id}`);
    const statusElement = document.getElementById(`status-${id}`);

    if (!noteElement || !statusElement) return;

    const note = noteElement.value;
    const status = statusElement.value;

    fetch(`/api/wishlist/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ note: note, status: status })
    })
        .then(res => {
            if (!res.ok) throw new Error('Could not update changes');
            alert('Changes saved successfully!');
        })
        .catch(err => alert(err.message));
}

// Delete from the wishlist
function deleteItem(id) {
    if (!confirm('Remove this destination?')) return;

    fetch(`/api/wishlist/${id}`, { method: 'DELETE' })
        .then(res => {
            if (!res.ok) throw new Error('Could not delete item');
            const row = document.getElementById(`row-${id}`);
            if (row) row.remove();

            //  If table is empty 

            const tbody = document.getElementById("wishlist-table-body");
            if (tbody && tbody.children.length === 0) {
                tbody.innerHTML = `
                <tr>
                    <td colspan="4" style="padding: 48px; text-align: center; color: #94a3b8; font-size: 15px;">
                      empty
                    </td>
                </tr>`;
            }
        })
        .catch(err => alert(err.message));
}
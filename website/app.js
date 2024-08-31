document.getElementById("addButton").addEventListener("click", function () {
	sendRequest("http://localhost:8080/add");
});

document
	.getElementById("subtractButton")
	.addEventListener("click", function () {
		sendRequest("http://localhost:8080/subtract");
	});

function sendRequest(endpoint) {
	const a = document.getElementById("number1").value || "0";
	const b = document.getElementById("number2").value || "0";

	const formData = new URLSearchParams();
	formData.append("a", a);
	formData.append("b", b);

	fetch(endpoint, {
		method: "POST",
		headers: {
			"Content-Type": "application/x-www-form-urlencoded",
		},
		body: formData,
	})
		.then((response) => {
			if (!response.ok) {
				return response.text().then((text) => {
					throw new Error(text);
				});
			}
			return response.json();
		})
		.then((data) => {
			document.getElementById("result").innerText = data.result;
		})
		.catch((error) => {
			document.getElementById("result").innerText =
				"Error: " + error.message;
			console.error("Error:", error);
		});
}

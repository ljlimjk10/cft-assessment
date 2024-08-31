document.getElementById("addButton").addEventListener("click", function () {
	sendRequest("http://localhost:8080/add");
});

document
	.getElementById("subtractButton")
	.addEventListener("click", function () {
		sendRequest("http://localhost:8080/subtract");
	});

function sendRequest(endpoint) {
	const num1 = document.getElementById("number1").value || "0";
	const num2 = document.getElementById("number2").value || "0";

	const formData = new URLSearchParams();
	formData.append("num1", num1);
	formData.append("num2", num2);

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

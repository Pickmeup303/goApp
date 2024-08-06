document.addEventListener('DOMContentLoaded', function() {
    const fileInput = document.getElementById('file_video');
    if (fileInput) {
        fileInput.addEventListener('change', function(e) {
            const formData = new FormData();
            formData.append('videoCapacity', fileInput.files[0]);

            fetch('/information_capacity', {
                method: 'POST',
                body: formData
            })
                .then(response => response.json())
                .then(data => {
                    if (data.validationError){
                        if (data.validationError.File){
                            alert(data.validationError.File)
                            document.getElementById('maxCapacity').innerText = "";
                        }
                    }else if (data.capacity){
                        document.getElementById('maxCapacity').innerText = data.capacity;
                    }
                })
                .catch(error => {
                    console.error('Error:', error);
                    alert('An error occurred while processing the request.');
                });
        });
    }
});

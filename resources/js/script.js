function put(data) {
    let xhr = new XMLHttpRequest();

    xhr.open("PUT", "/blog");
    
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(data));
}

function CreateBlog() {
    let caption = document.querySelector("#Caption");
    let content = document.querySelector("#Content");

    let data = {
        Caption: caption.value,
        Content: content.value,
    };

    put(data);
}
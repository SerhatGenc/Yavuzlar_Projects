<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dosya Yukleme </title>
</head>
<body>
    <h1>Dosya Yukleme</h1>
    <form method="GET" action="">
        <label for="file">Dosya adı:</label>
        <input type="text" id="file" name="file">
        <button type="submit">Yukle</button>
    </form>

    <?php
    if (isset($_GET['file'])) {

        $file = $_GET['file'];

        $filepath = "includes/" . $file;

        if (file_exists($filepath)) {
            include($filepath);
        } else {
            echo "Dosya bulunamadı.";
        }
    }
    ?>
</body>
</html>

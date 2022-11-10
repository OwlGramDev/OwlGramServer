<?php
header('Content-Type: image/jpeg');
$folder = dirname(__FILE__);
$data = file_get_contents($folder . '/apk_icon.jpg');
$banner = imagecreatefromstring($data);
$font_size = 65;
$font = $folder . '/GoogleSans-Bold.ttf';
$text = $argv[1];
$text_box = imagettfbbox($font_size, 0, $font, $text);
$text_width = $text_box[2]-$text_box[0];
$color = imagecolorallocate($banner, 255, 255, 255);
imagettftext($banner, $font_size, 0, (320 / 2) - ($text_width / 2), 250, $color, $font, $text);
ob_start();
imagejpeg($banner, null, 100);
$data = ob_get_clean();
file_put_contents($folder . '/output.jpg', $data);

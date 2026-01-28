
<?php
/**
 * Generate a TOTP token (RFC 6238) without external libraries
 */



function generateBase32Secret($length = 32) {
    $alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
    $secret = '';
    for ($i = 0; $i < $length; $i++) {
        $secret .= $alphabet[random_int(0, strlen($alphabet) - 1)];
    }
    return $secret;
}


// Base32 decode function
function base32Decode($secret) {
    $alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
    $secret = strtoupper($secret);
    $binaryString = '';
    for ($i = 0; $i < strlen($secret); $i++) {
        $char = $secret[$i];
        $pos = strpos($alphabet, $char);
        if ($pos === false) continue;
        $binaryString .= str_pad(decbin($pos), 5, '0', STR_PAD_LEFT);
    }
    $chunks = str_split($binaryString, 8);
    $result = '';
    foreach ($chunks as $chunk) {
        if (strlen($chunk) === 8) {
            $result .= chr(bindec($chunk));
        }
    }
    return $result;
}

// Generate TOTP
function generateTOTP($secret, $timeStep = 30, $digits = 6) {
    $key = base32Decode($secret);
    $counter = floor(time() / $timeStep);
    $binaryCounter = pack('N*', 0) . pack('N*', $counter);
    $hash = hash_hmac('sha1', $binaryCounter, $key, true);
    $offset = ord(substr($hash, -1)) & 0x0F;
    $truncatedHash = unpack('N', substr($hash, $offset, 4))[1] & 0x7FFFFFFF;
    $otp = $truncatedHash % pow(10, $digits);
    return str_pad($otp, $digits, '0', STR_PAD_LEFT);
}

// Example usage
// $secret = 'JBSWY3DPEHPK3PXP'; // Example Base32 secret
$password = "P@ssw0rd";
$secret = generateBase32Secret();
$token = generateTOTP($secret, 15 , 6);
echo "password: " . password_hash($password, PASSWORD_DEFAULT) . '<br/>';
echo "secret: " . $secret . '<br/>';
echo "TOTP Token: " . $token . '<br/>';
?>

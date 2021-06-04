const crypto = require('crypto')

const hash = crypto.createCipheriv('aes-192-ecb', Buffer.from(ENCRYPTION_KEY), iv);


const md5hash = crypto
                .createHash("md5")
                .update("Man oh man do I love node!")
                .digest("hex");

console.log('We don\'t want developers to use bad hash');
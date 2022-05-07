import * as ecies from '@walletconnect/ecies-25519';
import { Chacha20 } from "ts-chacha20";
import { createSHA3, createHMAC } from 'hash-wasm';

const oxyd = {
  utils: {
    u8ToBase64 (u8) {
      return btoa(String.fromCharCode.apply(null, u8));
    },
    base64ToU8 (str) {
      return atob(str).split('').map(function (c) { return c.charCodeAt(0); });
    },
  },

  async sha3(data) {
    let h = await createSHA3(256)
    h.update(data)
    return h.digest('binary')
  },

  async hmac(key, data) {
    const h = createSHA3(224)
    const m = await createHMAC(h, key)

    m.init()
    m.update(data)
    return m.digest('binary')
  },

  async chacha20_encrypt(key, data) {
    key = await this.sha3(key)

    const nonce = await ecies.randomBytes(12)
    const enc = new Chacha20(key, nonce).encrypt(data)

    let res = new Uint8Array(enc.length + nonce.length)
    res.set(enc)
    res.set(nonce, enc.length)

    return res
  },

  async chacha20_decrypt(key, data) {
    key = await this.sha3(key)

    const nonce = data.slice(-12)

    return new Chacha20(key, nonce).decrypt(data.slice(0, data.length-12))
  },

  async generate(password) {
    const keys = ecies.generateKeyPair(ecies.randomBytes(32))

    let pvk = keys.privateKey

    console.log("d1:", pvk)

    if (password) {
      pvk = await this.chacha20_encrypt(password, keys.privateKey)
    }

    return {
      public_key: keys.publicKey,
      private_key: pvk
    }
  }
}

export default oxyd
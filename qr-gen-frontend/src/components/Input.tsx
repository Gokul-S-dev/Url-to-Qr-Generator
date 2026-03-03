import React, { useState } from 'react'
import axios from 'axios'

const Input = () => {
  const [url, setUrl] = useState('')
  const [qrCode, setQrCode] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const generateQRCode = async () => {
    if (!url.trim()) {
      setError('Please enter a URL')
      return
    }

    setLoading(true)
    setError('')
    setQrCode('')

    try {
      const response = await axios.get(`http://localhost:8080/qr`, {
        params: { url: url },
        responseType: 'blob'
      })

      const imageUrl = URL.createObjectURL(response.data)
      setQrCode(imageUrl)
    } catch (err) {
      setError('Failed to generate QR code. Make sure the backend is running.')
      console.error('Error generating QR code:', err)
    } finally {
      setLoading(false)
    }
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      generateQRCode()
    }
  }

  return (
    <div className="qr-container">
      <h1>QR Code Generator</h1>
      <p className="subtitle">Enter a URL to generate a QR code</p>
      
      <div className="input-group">
        <input 
          type="text" 
          placeholder='Enter URL here (e.g., https://example.com)' 
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          onKeyPress={handleKeyPress}
          className="url-input"
        />
        <button 
          onClick={generateQRCode} 
          disabled={loading}
          className="generate-btn"
        >
          {loading ? 'Generating...' : 'Generate QR Code'}
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}

      {qrCode && (
        <div className="qr-result">
          <h3>Your QR Code:</h3>
          <img src={qrCode} alt="QR Code" className="qr-image" />
          <a 
            href={qrCode} 
            download="qrcode.png" 
            className="download-btn"
          >
            Download QR Code
          </a>
        </div>
      )}
    </div>
  )
}

export default Input
import { useEffect, useState } from "react";
import axios from "axios";

export default function WebPageAnalyzer() {
  const [url, setUrl] = useState("");
  const [urls, setUrls] = useState([]);
  const [selectedUrl, setSelectedUrl] = useState("");
  const [analysisResult, setAnalysisResult] = useState(null);
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [backendDown, setBackendDown] = useState(false);

  useEffect(() => {
    fetchUrls();
  }, []);

  const fetchUrls = async () => {
    try {
      const response = await axios.get("http://localhost:8080/urls");
      setUrls(response.data.urls || []); // Ensure urls is always an array
      setBackendDown(false);
    } catch (error) {
      console.error("Error fetching URLs:", error);
      setUrls([]); // Set an empty array in case of an error
      setBackendDown(true);
    }
  };

  const submitUrlForAnalysis = async () => {
    if (!url) return;
    setLoading(true);
    setErrorMessage("");
    setSuccessMessage("");
    try {
      const response = await axios.post("http://localhost:8080/analyze", new URLSearchParams({ url }));
      if (response.status === 202) {
        setSuccessMessage("Successfully submitted for analysis.");
        fetchUrls();
      } else {
        setErrorMessage("Unexpected response from the server.");
      }
    } catch (error) {
      console.error("Error submitting URL:", error);
      setErrorMessage("Failed to submit for analysis. Please try again.");
    }
    setLoading(false);
  };

  const fetchAnalysisResult = async () => {
    if (!selectedUrl) return;
    setLoading(true);
    setErrorMessage("");
    try {
      const response = await axios.get("http://localhost:8080/status", { params: { url: selectedUrl } });
      setAnalysisResult(response.data);
    } catch (error) {
      console.error("Error fetching analysis result:", error);
      setErrorMessage("Failed to fetch analysis result. Please try again.");
    }
    setLoading(false);
  };

  return (
    <div className="container mx-auto p-4">
      {backendDown && <div className="bg-red-500 text-white p-2 rounded mb-4">Backend is not reachable.</div>}
      {successMessage && <div className="bg-green-500 text-white p-2 rounded mb-4">{successMessage}</div>}
      {errorMessage && <div className="bg-red-500 text-white p-2 rounded mb-4">{errorMessage}</div>}

      <h1 className="text-xl font-bold mb-4">Web Page Analyzer</h1>

      <div className="mb-4">
        <input
          type="text"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="Enter URL"
          className="border p-2 rounded w-full"
        />
        <button
          onClick={submitUrlForAnalysis}
          className="bg-blue-500 text-white px-4 py-2 rounded mt-2"
          disabled={loading}
        >
          {loading ? "Submitting..." : "Submit for Analysis"}
        </button>
      </div>

      <div className="mb-4">
        <select
          value={selectedUrl}
          onChange={(e) => setSelectedUrl(e.target.value)}
          className="border p-2 rounded w-full"
        >
          <option value="">Select a URL</option>
          {urls.map((url, index) => (
            <option key={index} value={url}>{url}</option>
          ))}
        </select>
        <button
          onClick={fetchAnalysisResult}
          className="bg-green-500 text-white px-4 py-2 rounded mt-2"
          disabled={loading}
        >
          {loading ? "Fetching..." : "Get Analysis Result"}
        </button>
      </div>

      {analysisResult && (
        <div className="border p-4 rounded mt-4">
          <h2 className="text-lg font-bold">Analysis Result</h2>
          <pre className="bg-gray-100 p-2 rounded mt-2">{JSON.stringify(analysisResult, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

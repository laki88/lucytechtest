import { useEffect, useState, useRef } from "react";
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
  const [polling, setPolling] = useState(false);
  const [showBackendDownError, setShowBackendDownError] = useState(true);
  const [showSuccessMessage, setShowSuccessMessage] = useState(true);
  const [showErrorMessage, setShowErrorMessage] = useState(true);
  const [showWaitModal, setShowWaitModal] = useState(false);
  const [cancelMessage, setCancelMessage] = useState(false);

  const intervalRef = useRef(null);
  const timeoutRef = useRef(null);

  useEffect(() => {
    return () => {
      clearIntervalsAndTimeouts();
    };
  }, []);

  useEffect(() => {
    fetchUrls();
  }, []);

  useEffect(() => {
    if (backendDown) {
      const timer = setTimeout(() => {
        setShowBackendDownError(false); 
      }, 3000);
      return () => clearTimeout(timer);
    }
  }, [backendDown]);

  useEffect(() => {
    if (successMessage) {
      const timer = setTimeout(() => {
        setShowSuccessMessage(false); 
      }, 3000);
      return () => clearTimeout(timer);
    }
  }, [successMessage]);

  useEffect(() => {
    if (errorMessage) {
      const timer = setTimeout(() => {
        setShowErrorMessage(false); 
      }, 3000);
      return () => clearTimeout(timer);
    }
  }, [errorMessage]);

  useEffect(() => {
    if (cancelMessage) {
      const timer = setTimeout(() => {
        setCancelMessage(false);
      }, 3000);
      return () => clearTimeout(timer);
    }
  }, [cancelMessage]);

  const clearIntervalsAndTimeouts = () => {
    if (intervalRef.current) clearInterval(intervalRef.current);
    if (timeoutRef.current) clearTimeout(timeoutRef.current);
  };

  const fetchUrls = async () => {
    try {
      const response = await axios.get("http://localhost:8080/urls");
      setUrls(response.data.urls || []);
      setBackendDown(false);
    } catch (error) {
      console.error("Error fetching URLs:", error);
      setUrls([]);
      setBackendDown(true);
    }
  };

  const submitUrlForAnalysis = async () => {
    if (!url) return;
    setLoading(true);
    setErrorMessage("");
    setSuccessMessage("");
    setCancelMessage(false);
    setAnalysisResult(null);
    try {
      const response = await axios.post("http://localhost:8080/analyze", { url });
      if (response.status === 202) {
        setSuccessMessage("Successfully submitted for analysis.");
        setShowSuccessMessage(true);
        fetchUrls();
        startPolling(url);
      } else {
        setErrorMessage("Unexpected response from the server.");
      }
    } catch (error) {
      console.error("Error submitting URL:", error);
      setErrorMessage("Failed to submit for analysis. Because of : " + error.response?.data?.error);
      setLoading(false);
    }
  };

  const fetchAnalysisResult = async (fetchUrl,  isHistoryFetch = false) => {
    setCancelMessage(false);
    setLoading(true);
    setErrorMessage("");
    try {
      const response = await axios.get("http://localhost:8080/status", { 
        params: { url: fetchUrl || selectedUrl } 
      });
      if (response.data.Status === "Completed") {
        setAnalysisResult(response.data);
        setPolling(false);
        setLoading(false);
        clearIntervalsAndTimeouts();
        setShowWaitModal(false);
      } else if (isHistoryFetch) { 
        setAnalysisResult(response.data);
        setPolling(false);
        setLoading(false);
        clearIntervalsAndTimeouts();
        setShowWaitModal(false);
      } 
    } catch (error) {
      console.error("Error fetching analysis result:", error);
      setErrorMessage("Failed to fetch analysis result. Please try again.");
      setLoading(false);
    }  
  };

  const startPolling = (pollUrl) => {
    setPolling(true);
    clearIntervalsAndTimeouts();
    
    timeoutRef.current = setTimeout(() => {
      setShowWaitModal(true);
    }, 30000);
    
    intervalRef.current = setInterval(() => {
      fetchAnalysisResult(pollUrl);
    }, 5000);
  };

  const handleWaitModalOk = () => {
    setShowWaitModal(false);
  };

  const handleWaitModalCancel = () => {
    clearIntervalsAndTimeouts();
    setPolling(false);
    setLoading(false);
    setShowWaitModal(false);
    setCancelMessage(true);
  };

  return (
    <div className="container mx-auto p-4">
      {backendDown && showBackendDownError && (
        <div className="bg-red-500 text-white p-2 rounded mb-4">
          Backend is not reachable.
        </div>
      )}
      {successMessage && showSuccessMessage && (
        <div className="bg-green-500 text-white p-2 rounded mb-4">
          {successMessage}
        </div>
      )}
      {errorMessage && showErrorMessage && (
        <div className="bg-red-500 text-white p-2 rounded mb-4">
          {errorMessage}
        </div>
      )}
      {cancelMessage && (
        <div className="bg-yellow-500 text-white p-2 rounded mb-4">
          You may access the analysis result in history dropdown.
        </div>
      )}

      {showWaitModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded-lg max-w-md">
            <p className="mb-4">
              It seems like you have submitted a large web page for analysis. 
              This may take some time to finish. Would you like to wait?
            </p>
            <div className="flex justify-end space-x-3">
              <button
                onClick={handleWaitModalOk}
                className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
              >
                Continue Waiting
              </button>
              <button
                onClick={handleWaitModalCancel}
                className="bg-gray-500 text-white px-4 py-2 rounded hover:bg-gray-600"
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}

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
          Submit for Analysis
        </button>
      </div>
      
      <h2 className="text-lg font-bold mt-4">History</h2>
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
          onClick={() => {
            setCancelMessage(false);
            fetchAnalysisResult(selectedUrl, true);
          }}
          className="bg-green-500 text-white px-4 py-2 rounded mt-2"
          disabled={loading}
        >
          Get Analysis Result
        </button>
      </div>

      {loading && (
        <div className="flex justify-center my-4">
          <div className="animate-spin rounded-full h-8 w-8 border-t-4 border-blue-500 border-solid"></div>
        </div>
      )}

      {analysisResult && (
        <div className="border p-4 rounded mt-4 shadow-lg">
          <h2 className="text-lg font-bold">Analysis Result</h2>
          <table className="w-full border-collapse border border-gray-300 mt-2">
            <thead>
              <tr className="bg-gray-200 text-left">
                <th className="p-3 border border-gray-400">Key</th>
                <th className="p-3 border border-gray-400">Value</th>
              </tr>
            </thead>
            <tbody>
              {Object.entries(analysisResult).map(([key, value]) => (
                <tr key={key} className="hover:bg-gray-100">
                  <td className="p-3 border border-gray-400 font-medium">{key}</td>
                  <td className="p-3 border border-gray-400">
                    {JSON.stringify(value).replace(/\"/g, "")}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
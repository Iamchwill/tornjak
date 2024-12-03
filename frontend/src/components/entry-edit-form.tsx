import React, { useState, useEffect } from 'react';

const EditServerForm = ({ server, onSave, onClose, globalSelectorInfo, globalAgentsList, globalEntryExpiryTime, globalWorkloadSelectorInfo, globalAgentsWorkLoadAttestorInfo, globalDebugServerInfo }) => {
  const [name, setName] = useState(server.name);
  const [ip, setIp] = useState(server.ip);
  const [status, setStatus] = useState(server.status);
  const [selectorInfo, setSelectorInfo] = useState(globalSelectorInfo);
  const [agentsList, setAgentsList] = useState(globalAgentsList);

  useEffect(() => {
    setName(server.name);
    setIp(server.ip);
    setStatus(server.status);
  }, [server]);

  const handleSubmit = (e) => {
    e.preventDefault();

    const updatedServer = { ...server, name, ip, status, selectorInfo, agentsList };
    onSave(updatedServer);  // Call the parent onSave method with the updated data
  };

  return (
    <div className="edit-server-form">
      <h3>Edit Server</h3>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name:</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div>
          <label>IP Address:</label>
          <input
            type="text"
            value={ip}
            onChange={(e) => setIp(e.target.value)}
          />
        </div>
        <div>
          <label>Status:</label>
          <input
            type="text"
            value={status}
            onChange={(e) => setStatus(e.target.value)}
          />
        </div>

        <div>
          <label>Selector Info:</label>
          <input
            type="text"
            value={JSON.stringify(selectorInfo)} 
            onChange={(e) => setSelectorInfo(JSON.parse(e.target.value))}
          />
        </div>

        <div>
          <label>Agents List:</label>
          <input
            type="text"
            value={JSON.stringify(agentsList)}
            onChange={(e) => setAgentsList(JSON.parse(e.target.value))}
          />
        </div>

        <button type="submit">Save</button>
        <button type="button" onClick={onClose}>Cancel</button>
      </form>
    </div>
  );
};

export default EditServerForm;

export async function GetQR(clientId){
    const resp = await fetch(`/api/qr?clientId=${clientId}`,{
        method:"GET",
        
    })

    return await resp.json()
}

export async function GetSession(sessionId){
    const resp = await fetch(`/api/sessions/${sessionId}`,{
        method:"GET",
        
    })
    if (resp.status === 204) {
        return {connected:false}
    }
    if (resp.status ===200){
        const user = await resp.json()
        return {connected:true, user:user}
    }
    const message = await resp.json()
    throw new Error(`Server returns error. Status ${resp.status} Error ${message}`)
}

export async function GenerateProof(data,sessionId){
    const resp = await fetch(`/api/sessions/${sessionId}`,{
        method:"POST",
        body: JSON.stringify(data)        
    })
    return await resp.text()

}

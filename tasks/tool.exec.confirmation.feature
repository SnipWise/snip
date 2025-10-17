Lorsque commence le flow de la completion en streaming 
(avec la fonction DefineStreamingChatFlow dans /chatflow/streaming.go)

Je voudrais, qu'il y ait une pause (le statut de l'operation serait en mode pending)
Que le flow envoie au client web un message json du type: 

{
    "message": "tool detected",
    "status": "pending" 
}

et le flow ne reprendrait uniquement si on a appelé 
le endpoint /operation/validate ou /operation/cancel
qui mettent à jour le statut de l'operation

Avec Genkit il est possible d'interrompre et reprendre une generation comme expliquer ici:
https://genkit.dev/docs/interrupts/?lang=go

- Ajoute cette pause et attente de confirmation à la ligne 62 de /chatflow/streaming.go
- Cree les 2 endpoints /operation/validate et /operation/cancel qui permettent de debloquer la pause
- Cree 2 fichiers bash pour appeler chacun des endpoints ci dessus

Il faudra que tu testes que cela fonctionne: pour lancer une stream completion tu peux utiliser le script ./stream.sh 

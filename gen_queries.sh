set -o errexit
set -o pipefail

#ignite scaffold query card cardId:uint --response card:CardWithImage
# ---ignite scaffold query card_content cardId:uint --response card:Card
#ignite scaffold query user address:string --response user:User
#ignite scaffold query cards owner --response cardIds:uints
ignite scaffold query match matchId --response match:Match
ignite scaffold query set setId --response match:SetWithArtwork

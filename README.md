# feedthebaby
A bot built to send Discord notifications when baby formula becomes available in store.

*Note: Please do not use this to scalp formula. This was built solely for purchasing formula for USE, not for resale. Don't be selfish.*

## What is this world coming to?

I built this bot in the midst of the "great formula shortage" so I could become notified when formula was back in stock so that I could then purchase it for our baby. Now that we have completely stocked up on as much formula as necessary, I am sharing it so others have the opportunity to do the same. As said above, please do not use this for scalping purposes!

## Do I need to modify anything?

Make sure you modify the constant in `main.go` (`DISCORD_WEBHOOK`) to match up with your Discord webhook integration.  If you need help finding/setting this up see: https://hookdeck.com/webhooks/platforms/how-to-get-started-with-discord-webhooks#what-do-webhooks-do-in-discord

## How do I modify it for other types of formula?

I'll leave figuring out this part up to you, but it is fairly straight forward. Feel free to contribute to the project if you make modifications that you think may help others!

## Usage

1. Install Go
2. git clone https://github.com/un4gi/feedthebaby.git
3. cd feedthebaby
4. *Modify the DISCORD_WEBHOOK constant in main.go and adjust the code as necessary for other types of formula.*
5. go build .
6. ./feedthebaby

*Note: I recommend scheduling a cron job/task to run this every minute to ensure it does what it was intended to do.*

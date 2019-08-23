let index = {
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            index.ojimessage();
        })
    },
    ojimessage: function() {
        // Create message
        let message = {"name": "ojimessage"};
        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            // View
            document.getElementById("ojisan_message").innerText = message.payload;
        })
    }
};
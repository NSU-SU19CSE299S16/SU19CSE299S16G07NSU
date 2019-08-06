(function () {
    var pusher = new Pusher('PUSHER_APP_KEY', {
        authEndpoint: '/pusher/auth',
        cluster: 'PUSHER_APP_CLUSTER',
        encrypted: true
    });

    let chat = {
        name: undefined,
        email: undefined,
        endUserName: undefined,
        currentRoom: undefined,
        currentChannel: undefined,
        subscribedChannels: [],
        subscribedUsers: []
    }

    var publicChannel = pusher.subscribe('update');

    const chatBody = $(document)
    const chatRoomsList = $('#rooms')
    const chatReplyMessage = $('#replyMessage')

    const helpers = {
        clearChatMessages: () => {
            $('#chat-msgs').html('')
        },

        displayChatMessage: (message) => {
            if (message.email === chat.email) {
                $('#chat-msgs').prepend(
                    `<tr>
                            <td>
                                <div class="sender">${message.sender} @ <span class="date">${message.createdAt}</span></div>
                                <div class="message">${message.text}</div>
                            </td>
                        </tr>`
                )
            }
        },

        loadChatRoom: evt => {
            chat.currentRoom = evt.target.dataset.roomId
            chat.currentChannel = evt.target.dataset.channelId
            chat.endUserName =  evt.target.dataset.userName
            if (chat.currentRoom !== undefined) {
                $('.response').show()
                $('#room-title').text('Write a message to ' + evt.target.dataset.userName+ '.')
            }

            evt.preventDefault()
            helpers.clearChatMessages()
        },

        replyMessage: evt => {
            evt.preventDefault()

            let createdAt = new Date().toLocaleString()
            let message = $('#replyMessage input').val().trim()
            let event = 'client-' + chat.currentRoom

            chat.subscribedChannels[chat.currentChannel].trigger(event, {
                'sender': chat.name,
                'email': chat.currentRoom,
                'text': message,
                'createdAt': createdAt
            });

            $('#chat-msgs').prepend(
                `<tr>
                        <td>
                            <div class="sender">
                                ${chat.name} @ <span class="date">${createdAt}</span>
                            </div>
                            <div class="message">${message}</div>
                        </td>
                    </tr>`
            )

            $('#replyMessage input').val('')
        },

        LogIntoChatSession: function (evt) {
            const name  = $('#fullname').val().trim()
            const email = $('#email').val().trim().toLowerCase()

            chat.name = name;
            chat.email = email;

            chatBody.find('#loginScreenForm input, #loginScreenForm button').attr('disabled', true)

            let validName = (name !== '' && name.length >= 3)
            let validEmail = (email !== '' && email.length >= 5)

            if (validName && validEmail) {
                axios.post('/new/user', {name, email}).then(res => {
                    chatBody.find('#registerScreen').css("display", "none");
                    chatBody.find('#main').css("display", "block");

                    chat.myChannel = pusher.subscribe('private-' + res.data.email)
                    chat.myChannel.bind('client-' + chat.email, data => {
                        helpers.displayChatMessage(data)
                    })
                })
            } else {
                alert('Enter a valid name and email.')
            }

            evt.preventDefault()
        }
    }


    publicChannel.bind('new-user', function(data) {
        if (data.email != chat.email){
            chat.subscribedChannels.push(pusher.subscribe('private-' + data.email));
            chat.subscribedUsers.push(data);

            $('#rooms').html("");

            chat.subscribedUsers.forEach((user, index) => {
                $('#rooms').append(
                    `<li class="nav-item"><a data-room-id="${user.email}" data-user-name="${user.name}" data-channel-id="${index}" class="nav-link" href="#">${user.name}</a></li>`
                )
            })
        }
    })

    chatReplyMessage.on('submit', helpers.replyMessage)
    chatRoomsList.on('click', 'li', helpers.loadChatRoom)
    chatBody.find('#loginScreenForm').on('submit', helpers.LogIntoChatSession)
}());
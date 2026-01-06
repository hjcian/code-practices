CreateUser(userName, preparePublisher(userName));
const userName = "MaxCian";


function CreateUser(userName, eventPublisher = null) {
    console.log(`User ${userName} created in datastore.`);
    if (eventPublisher != null) {
        eventPublisher();
    }
}
function preparePublisher(name) {
    // This is a closure that captures the userName variable
    return ()=> console.log(`Event: User ${userName} created`);
}

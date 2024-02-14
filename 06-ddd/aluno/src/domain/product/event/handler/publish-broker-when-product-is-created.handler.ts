import EventHandlerInterface from "../../../@shared/event/event-handler.interface";
import ProductCreatedEvent from "../product-created.event";

export default class PublishBrokerWhenProductIsCreatedHandler implements EventHandlerInterface<ProductCreatedEvent> {
    
    handle(event: ProductCreatedEvent): void {
        console.log(`Sending to broker..... Eventdata:`, event);
    }

}
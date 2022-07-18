import 'package:grpc/grpc.dart';
import 'package:event_mate/pb/profile-service.pbgrpc.dart';
import 'package:event_mate/pb/profile-service.pb.dart';

//TODO refeactor et , class a yada fonk a cevir bu sekilde cmd den dart client.dart komutu ile calisiyor
Future<void> main() async {
  final channel = ClientChannel(
    'localhost',
    port: 50051,
    options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
  );

  final client = ProfileServiceClient(channel);

  try {
    var response = await client.getUser(GetUserEventRequest()..userId = 1);
    print('Profile client received: ${response.user.name}');
  } catch (e) {
    print('Caught error: $e');
  }
  await channel.shutdown();
}
import 'package:event_mate/firebase_options.dart';
import 'package:event_mate/modules/screens/home_screen.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter/material.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:flutter/foundation.dart' show kIsWeb;

@pragma('vm:entry-point')
Future<void> _firebaseMessagingBackgroundHandler(RemoteMessage message) async {
  // If you're going to use other Firebase services in the background, such as Firestore,
  // make sure you call `initializeApp` before using other Firebase services.

  print("Handling a background message: ${message.messageId}");
}

class Authentication {
  static SnackBar customSnackBar({required String content}) {
    return SnackBar(
      backgroundColor: Colors.black,
      content: Text(
        content,
        style: TextStyle(color: Colors.redAccent, letterSpacing: 0.5),
      ),
    );
  }

  static Future<FirebaseApp> initializeFirebase({
    required BuildContext context,
  }) async {
    WidgetsFlutterBinding.ensureInitialized();
    FirebaseApp firebaseApp = await Firebase.initializeApp(
        options: DefaultFirebaseOptions.currentPlatform);

    User? user = FirebaseAuth.instance.currentUser;

    if (user != null) {
      Navigator.of(context).pushReplacement(
        MaterialPageRoute(
          builder: (context) => HomeScreen(
            firebaseUser: user,
          ),
        ),
      );
    }
    final fcmToken = await FirebaseMessaging.instance.getToken();
    print("token$fcmToken");
    FirebaseMessaging.instance.onTokenRefresh.listen((fcmToken) {
      // TODO: If necessary send token to application server.

      // Note: This callback is fired at each app startup and whenever a new
      // token is generated.
      print("token$fcmToken");
    }).onError((err) {
      // Error getting token.
      print("err$err");
    });

    FirebaseMessaging.onBackgroundMessage(_firebaseMessagingBackgroundHandler);

    FirebaseMessaging messaging = FirebaseMessaging.instance;

    NotificationSettings settings = await messaging.requestPermission(
      alert: true,
      badge: true,
      provisional: false,
      sound: true,
    );

    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      print("on message");
    });

    return firebaseApp;
  }

  static Future<User?> signInWithGoogle({required BuildContext context}) async {
    FirebaseAuth auth = FirebaseAuth.instance;
    User? user;

    if (kIsWeb) {
      GoogleAuthProvider authProvider = GoogleAuthProvider();

      try {
        final UserCredential userCredential =
            await auth.signInWithPopup(authProvider);

        user = userCredential.user;
      } catch (e) {
        print(e);
      }
    } else {
      final GoogleSignIn googleSignIn = GoogleSignIn();

      final GoogleSignInAccount? googleSignInAccount =
          await googleSignIn.signIn();

      if (googleSignInAccount != null) {
        final GoogleSignInAuthentication googleSignInAuthentication =
            await googleSignInAccount.authentication;

        final AuthCredential credential = GoogleAuthProvider.credential(
          accessToken: googleSignInAuthentication.accessToken,
          idToken: googleSignInAuthentication.idToken,
        );

        try {
          final UserCredential userCredential =
              await auth.signInWithCredential(credential);

          user = userCredential.user;
        } on FirebaseAuthException catch (e) {
          if (e.code == 'account-exists-with-different-credential') {
            ScaffoldMessenger.of(context).showSnackBar(
              Authentication.customSnackBar(
                content:
                    'The account already exists with a different credential',
              ),
            );
          } else if (e.code == 'invalid-credential') {
            ScaffoldMessenger.of(context).showSnackBar(
              Authentication.customSnackBar(
                content:
                    'Error occurred while accessing credentials. Try again.',
              ),
            );
          }
        } catch (e) {
          ScaffoldMessenger.of(context).showSnackBar(
            Authentication.customSnackBar(
              content: 'Error occurred using Google Sign In. Try again.',
            ),
          );
        }
      }
    }

    return user;
  }

  static Future<void> signOut({required BuildContext context}) async {
    final GoogleSignIn googleSignIn = GoogleSignIn();
    try {
      await FirebaseAuth.instance.signOut();
      googleSignIn.signOut();
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(Authentication.customSnackBar(
          content: 'Error signing out. Try again.'));
    }
  }
}
